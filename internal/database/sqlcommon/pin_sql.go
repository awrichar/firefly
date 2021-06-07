// Copyright © 2021 Kaleido, Inc.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sqlcommon

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/kaleido-io/firefly/internal/i18n"
	"github.com/kaleido-io/firefly/internal/log"
	"github.com/kaleido-io/firefly/pkg/database"
	"github.com/kaleido-io/firefly/pkg/fftypes"
)

var (
	pinColumns = []string{
		"masked",
		"hash",
		"batch_id",
		"idx",
		"dispatched",
		"created",
	}
	pinFilterTypeMap = map[string]string{
		"batch": "batch_id",
		"index": "idx",
	}
)

func (s *SQLCommon) UpsertPin(ctx context.Context, pin *fftypes.Pin) (err error) {
	ctx, tx, autoCommit, err := s.beginOrUseTx(ctx)
	if err != nil {
		return err
	}
	defer s.rollbackTx(ctx, tx, autoCommit)

	pin.Sequence, err = s.insertTx(ctx, tx,
		sq.Insert("pins").
			Columns(pinColumns...).
			Values(
				pin.Masked,
				pin.Hash,
				pin.Batch,
				pin.Index,
				pin.Dispatched,
				pin.Created,
			),
	)
	if err == nil {
		s.postCommitEvent(tx, func() {
			s.callbacks.PinCreated(pin.Sequence)
		})
	} else {
		// Check it's not just that it already exsits (edge case, so we optimize for insert)
		pinRows, queryErr := s.queryTx(ctx, tx,
			sq.Select("masked", s.provider.SequenceField(""), "dispatched").
				From("pins").
				Where(sq.Eq{
					"hash":     pin.Hash,
					"batch_id": pin.Batch,
					"idx":      pin.Index,
				}),
		)
		existing := false
		if queryErr == nil {
			defer pinRows.Close()
			existing = pinRows.Next()
		}
		if !existing {
			// Something else went wrong - return the original error
			return err
		}
		err := pinRows.Scan(&pin.Masked, &pin.Sequence, &pin.Dispatched)
		if err != nil {
			return i18n.WrapError(ctx, err, i18n.MsgDBReadErr, "pins")
		}
		log.L(ctx).Debugf("Pin already existed")
		return nil
	}

	return s.commitTx(ctx, tx, autoCommit)
}

func (s *SQLCommon) pinResult(ctx context.Context, row *sql.Rows) (*fftypes.Pin, error) {
	pin := fftypes.Pin{}
	err := row.Scan(
		&pin.Masked,
		&pin.Hash,
		&pin.Batch,
		&pin.Index,
		&pin.Dispatched,
		&pin.Created,
		&pin.Sequence,
	)
	if err != nil {
		return nil, i18n.WrapError(ctx, err, i18n.MsgDBReadErr, "pins")
	}
	return &pin, nil
}

func (s *SQLCommon) GetPins(ctx context.Context, filter database.Filter) (message []*fftypes.Pin, err error) {

	cols := append([]string{}, pinColumns...)
	cols = append(cols, s.provider.SequenceField(""))
	query, err := s.filterSelect(ctx, "", sq.Select(cols...).From("pins"), filter, pinFilterTypeMap)
	if err != nil {
		return nil, err
	}

	rows, err := s.query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	pin := []*fftypes.Pin{}
	for rows.Next() {
		d, err := s.pinResult(ctx, rows)
		if err != nil {
			return nil, err
		}
		pin = append(pin, d)
	}

	return pin, err

}

func (s *SQLCommon) SetPinDispatched(ctx context.Context, sequence int64) (err error) {

	ctx, tx, autoCommit, err := s.beginOrUseTx(ctx)
	if err != nil {
		return err
	}
	defer s.rollbackTx(ctx, tx, autoCommit)

	err = s.updateTx(ctx, tx, sq.
		Update("pins").
		Set("dispatched", true).
		Where(sq.Eq{
			s.provider.SequenceField(""): sequence,
		}))
	if err != nil {
		return err
	}

	return s.commitTx(ctx, tx, autoCommit)
}

func (s *SQLCommon) DeletePin(ctx context.Context, sequence int64) (err error) {

	ctx, tx, autoCommit, err := s.beginOrUseTx(ctx)
	if err != nil {
		return err
	}
	defer s.rollbackTx(ctx, tx, autoCommit)

	err = s.deleteTx(ctx, tx, sq.Delete("pins").Where(sq.Eq{
		s.provider.SequenceField(""): sequence,
	}))
	if err != nil {
		return err
	}

	return s.commitTx(ctx, tx, autoCommit)
}
