
function trackHovers() {
    document.querySelectorAll('.track-hover').forEach(function(element){
        element.addEventListener('mouseenter', function(event){
            const data = {
                eventType: 'hover',
                element: event.target.tagName,
                id: event.target.id,
                classes: event.target.className,
                textContent: event.target.textContent,
                timestamp: new Date().toISOString()
            }
            sendData(data, 'mousehover')
        })
    })
}

export default trackHovers;