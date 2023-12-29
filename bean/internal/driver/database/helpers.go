package database

import (
	"fmt"
)

func dsn(builder *DSNBuilder) string {
	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v)/%v?tls=%v&interpolateParams=%v",
		builder.Username,
		builder.Password,
		builder.Host,
		builder.Name,
		builder.Tls,
		builder.Interpolate,
	)

	return dsn
}
