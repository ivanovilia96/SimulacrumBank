package client

import (
	"errors"
	"fmt"
	"strings"
)

// if arg is nil, then this field wont be changed
func (c Client) Update(currentMail, newMail, newFio string, newAge int8) error {
	var changeList []string
	if newMail != "" {
		changeList = append(changeList, fmt.Sprintf(`mail = '%s'`, newMail))
	}
	if newFio != "" {
		changeList = append(changeList, fmt.Sprintf("fio = '%s'", newFio))
	}
	if newAge != 0 {
		changeList = append(changeList, fmt.Sprintf("age = %d", newAge))
	}

	sqlStatement := fmt.Sprintf("UPDATE client SET %s WHERE mail = '%s'", strings.Join(changeList, ", "), currentMail)

	r, err := c.db.Exec(sqlStatement)
	if err != nil {
		return err
	}

	rowsAffected, err := r.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("we didnt find any clients with this email " + currentMail)
	}

	return nil
}
