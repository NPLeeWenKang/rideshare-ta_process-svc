package main

type Driver struct {
	Driver_Id    int    `json:"driver_id"`
	First_Name   string `json:"first_name"`
	Last_Name    string `json:"last_name"`
	Mobile_No    string `json:"mobile_no"`
	Email        string `json:"email"`
	Id_No        string `json:"id_no"`
	Car_No       string `json:"car_no"`
	Is_Available bool   `json:"is_available"`
}
