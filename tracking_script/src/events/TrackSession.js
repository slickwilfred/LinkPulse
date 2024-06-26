import sendData from "../services/apiService";

const sessionStart = Date.now();

function trackSession() {
    window.addEventListener('unload', function() {
        const sessionEnd = Date.now();
        const sessionDuration = sessionEnd - sessionStart;
        const data = {
            eventType: 'sessionEnd',
            duration: sessionDuration,
            start: new Date(sessionStart).toISOString(),
            end: new Date(sessionEnd).toISOString(),
            url: window.location.href
        }
        sendData(data, 'sessionDuration')
    })
}

export default trackSession;