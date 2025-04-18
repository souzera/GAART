import axios from "axios";
import { API_ROOT } from "./config";

export default async function login(username, password) {
    const response = await axios(`${API_ROOT}/login`, {
        method: 'POST',
        headers: {

            'Content-Type': 'application/json',
        },
        data: {
            login: username,
            senha: password,
        },
    }).catch((error) => {   
        console.error('Error:', error);
        return null;
    }
    );

    if (response && response.status === 200) {
        let data = response.data.data;
        document.cookie = `token=${data.token}; path=/; max-age=86400`;
    } else {
        return null;
    }

    return response
}
