import sendData from "../services/apiService";


/**
 * 
 */
function trackClicks () {
    document.addEventListener('click', function(event) {
        const data = {
            eventType: 'click',
            element: event.target.tagName,
            classes: event.target.className,
            id: event.target.id,
            textContent: event.target.textContent,
            timestamp: new Date().toISOString()
        }
        sendData(data, 'clicks')
    })
}

export default trackClicks;