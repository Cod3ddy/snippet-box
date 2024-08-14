package models

import (
	"testing"

	"github.com/Cod3ddy/snippet-box/internal/assert"
)

func TestIUserModelExists(t *testing.T) {
	if testing.Short() {
		t.Skip("models: skipping integration test")
	}
	// tdts 
	tests := []struct{
		name string
		userID int
		want bool
	}{
		{
			name: "Valid ID", 
			userID: 1,
			want: true,
		},
		{
            name:   "Zero ID",
            userID: 0,
            want:   false,
        },
        {
            name:   "Non-existent ID",
            userID: 2,
            want:   false,
        },
	}

	for _, tt := range tests{
		t.Run(tt.name, func(t *testing.T){
			// call new testdb() helper
			db := newTestDB(t)

			m := NewUserModel(db)

			exists, err := m.Exists(tt.userID)
			
			assert.Equal(t, exists, tt.want)
			assert.NilError(t, err)
		})
	}
}