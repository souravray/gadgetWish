/*
* @Author: souravray
* @Date:   2014-10-21 03:35:32
* @Last Modified by:   souravray
* @Last Modified time: 2014-10-21 03:36:22
 */

package models

import (
	"errors"
	"regexp"
)

var (
	EmailRegexp = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
)

var (
	ErrInvalidEmail                  = errors.New(`Invalid email`)
	ErrNotFilled                     = errors.New(`Blank email`)
	ErrorProductNameNotFilled        = errors.New(`Blank product name`)
	ErrorProductImgNotFilled         = errors.New(`Blank product image`)
	ErrorProductDescriptionNotFilled = errors.New(`Blank product description`)
	ErrorProductPriceNotFilled       = errors.New(`Blank product price`)
	ErrorProductPriceUnitNotFilled   = errors.New(`Blank product price-unit`)
)

func init() {

}
