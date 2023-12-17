package xerrors

import (
	errors2 "errors"
	"fmt"
	"github.com/pkg/errors"
)

type DbErr struct {
	msg string
	err error
}

func (e DbErr) Error() string {
	return fmt.Sprintf("DbErr(%s)", e.msg)
}

type DbErrInsert struct {
}

func workWrapErr() error {
	err := wrapErr()
	if errors2.As(err, &DbErr{msg: "bad insert", err: errors2.New("h")}) {
		fmt.Println("this is db err", err)
	}

	fmt.Println("workWrapErr done!")
	return nil
}

func wrapErr() error {
	return fmt.Errorf("%w", DbErr{
		msg: "bad insert",
		err: nil,
	})
}

func Work() {
	e := service()
	fmt.Printf("%v", e)
}

func service() error {
	err := middle()

	return err
}

func middle() error {
	err := daos()
	uid := 90
	//errors.New()
	err = errors.Wrap(err, fmt.Sprintf("middle failed uid(%+v)", uid))
	return err
}

func daos() error {

	err := errors.Wrap(errors.New("basic"), "daos failed")

	return err
}
