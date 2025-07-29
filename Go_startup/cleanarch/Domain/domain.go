package Domain


type Task struct {
	ID          string   
	Title       string    
	Description string    
	DueDate     string		
	Status      string    
	UserID      string    
}

type User struct {
	ID       string 
	Username string 
	Password string 
	Role     string 
}