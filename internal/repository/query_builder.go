package repository

import sq "github.com/Masterminds/squirrel"

var qbuilder = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
