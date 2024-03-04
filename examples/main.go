package main

import (
	"fmt"

	"github.com/broderick-westrope/goenums/examples/enums/auth"
)

// NOTE: normally this would look more like `go:generate goenums ./config/enums.json ./enums/` when you have installed the goenums package.
//go:generate go run ../cmd/goenums/ ./config/enums.json ./enums/

// Example usage of the generated enums and their extensions.
func main() {
	role := auth.Roles.ADMIN
	fmt.Printf("Role: %d, String: %s\n", role, role.String())
	// Role: 0, String: Admin

	status := auth.AccountStatusContainer{}.ACTIVE
	fmt.Printf("Status: %d, String: %s\n", status, status.String())
	// Status: 0, String: Active

	if role.CanAccessAdminPanel() {
		fmt.Println("Access granted to the admin panel.")
	}
	// Access granted to the admin panel.

	fmt.Println(status.StatusMessage())
	// Your account is active.
}
