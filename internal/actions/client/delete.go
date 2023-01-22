package client

import (
	"fmt"
)

// Delete - delete client with mail (PK)
func (c Client) Delete(mail string) error {
	добавить что бы при удалении клиента все его счета удалялись ( каскадное удаление )
	result, err := c.db.Exec("DELETE FROM client WHERE mail=$1;", mail)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Printf("deleted %d rows", rowsAffected)
	if rowsAffected == 0 {
		return ErrorNoDeletedRows
	}

	return nil
}
