package client

import "errors"

var ErrorNoDeletedRows = errors.New("there isn`t row with this email ")
