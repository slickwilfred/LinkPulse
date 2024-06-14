dotenv.config()

/**
 * 
 */
function sendData(data, path) {
    const API_URL = process.env.API_URL;
    fetch(`${API_URL}/${path}`, {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify(data),
        keepalive: true // Important for send-operations in unload events
    }).catch(error => {
        console.error('Failed to send data', error);
    })
}

export default sendData;