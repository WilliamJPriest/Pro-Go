package routes

import(
	"testing"
)

func Test_Accounts(t *testing.T){
	t.Run("Working Register Route",Test_Register_Working)
	t.Run("Working Login Route",Test_Login_Failed)
	t.Run("Failed Register Route",Test_Register_Failed)
	t.Run("Failed Login Route",Test_Login_Failed)
	

}

func Test_Register_Working(t *testing.T){
	
}

func Test_Login_Working(t *testing.T){
	
}

func Test_Register_Failed(t *testing.T){
	
}

func Test_Login_Failed(t *testing.T){
	
}