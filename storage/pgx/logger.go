package pgx

import (
	"context"

	"github.com/underbek/examples-go/logger"
	"github.com/underbek/examples-go/tracing"
	"go.opentelemetry.io/otel/trace"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type DBLogger struct {
	conn   Storage
	logger *logger.Logger
}

type TxLogger struct {
	tx     Transaction
	logger *logger.Logger
}

func NewWithLogger(ctx context.Context, dataSource string, logger *logger.Logger) (Storage, error) {
	db, err := New(ctx, dataSource)
	if err != nil {
		return nil, err
	}

	return &DBLogger{
		conn:   db,
		logger: logger.Named("storage"),
	}, nil

}

func (l *DBLogger) Exec(ctx context.Context, sql string, args ...interface{}) (commandTag pgconn.CommandTag, err error) {
	commandTag, err = l.conn.Exec(ctx, sql, args...)
	l.logger.WithCtx(ctx).With("sql", sql).With("arguments", args).WithError(err).Debug("Exec")

	return commandTag, err
}

func (l *DBLogger) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	ctx, span := tracing.StartCustomSpan(ctx, trace.SpanKindInternal, "pgx", "Query")
	defer span.End()

	rows, err := l.conn.Query(ctx, sql, args...)
	l.logger.WithCtx(ctx).With("sql", sql).With("arguments", args).WithError(err).Debug("Query")

	return rows, err
}

func (l *DBLogger) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	ctx, span := tracing.StartCustomSpan(ctx, trace.SpanKindInternal, "pgx", "QueryRow")
	defer span.End()

	row := l.conn.QueryRow(ctx, sql, args...)
	l.logger.WithCtx(ctx).With("sql", sql).With("arguments", args).Debug("QueryRow")

	return row
}

func (l *DBLogger) Close() {
	l.conn.Close()
	l.logger.Debug("Close")
}

// Begin returned transaction wrapper
func (l *DBLogger) Begin(ctx context.Context, opts *pgx.TxOptions) (Transaction, error) {
	tx, err := l.conn.Begin(ctx, opts)
	l.logger.WithCtx(ctx).WithError(err).With("options", opts).Debug("Begin")
	if err != nil {
		return nil, err
	}

	return &TxLogger{
		tx:     tx,
		logger: l.logger,
	}, nil
}

//
// Transaction methods below
//

func (l *TxLogger) Exec(ctx context.Context, sql string, args ...interface{}) (commandTag pgconn.CommandTag, err error) {
	commandTag, err = l.tx.Exec(ctx, sql, args...)
	l.logger.WithCtx(ctx).With("sql", sql).With("arguments", args).WithError(err).Debug("ExecTx")

	return commandTag, err
}

func (l *TxLogger) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	ctx, span := tracing.StartCustomSpan(ctx, trace.SpanKindInternal, "pgx", "QueryTx")
	defer span.End()

	rows, err := l.tx.Query(ctx, sql, args...)
	l.logger.WithCtx(ctx).With("sql", sql).With("arguments", args).WithError(err).Debug("QueryTx")

	return rows, err
}

func (l *TxLogger) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	ctx, span := tracing.StartCustomSpan(ctx, trace.SpanKindInternal, "pgx", "QueryRowTx")
	defer span.End()

	row := l.tx.QueryRow(ctx, sql, args...)
	l.logger.WithCtx(ctx).With("sql", sql).With("arguments", args).Debug("QueryRowTx")

	return row
}

func (l *TxLogger) Commit(ctx context.Context) error {
	err := l.tx.Commit(ctx)
	l.logger.WithCtx(ctx).WithError(err).Debug("Commit")

	return err
}

func (l *TxLogger) Rollback(ctx context.Context) error {
	err := l.tx.Rollback(ctx)
	l.logger.WithCtx(ctx).WithError(err).Debug("Rollback")

	return err
}