package messages

// GetMessagesResp devuelve los mensajes pre establecidos.
func GetMessagesResp() map[string]string {
	messages := make(map[string]string)

	//Declaracion de todos los mensajes que se quieran incorporar en todos los responses de la API
	//messages["IMPORTANTE"] = "Todos los nombres de las propiedades deben tener todos sus caracteres en 'minusculas'"

	return messages
}
