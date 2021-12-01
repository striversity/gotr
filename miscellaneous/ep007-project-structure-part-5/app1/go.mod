module github.com/striversity/misc007/app1

go 1.17

replace github.com/striversity/misc007/moda => ../moda

require (
	github.com/sirupsen/logrus v1.8.1
	github.com/striversity/misc007/moda v0.0.0-00010101000000-000000000000
)

require golang.org/x/sys v0.0.0-20191026070338-33540a1f6037 // indirect
