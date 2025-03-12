const GatewayURL = "http://localhost:8000/"
const NotesURL = GatewayURL + "notes/"
const UserURL = GatewayURL + "user/"

export const GetNoteURL = (id: number) => {
    return NotesURL + String(id);
}

export const RegisterURL = UserURL + "/register" 
export const LoginURL = UserURL + "/login" 
export const ChangeURL = UserURL + "/change" 