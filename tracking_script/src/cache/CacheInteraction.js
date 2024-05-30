import sendData from "../services/apiService";

let interactionCache = [];

function cacheInteraction(data) {
    interactionCache.push(data);

    // Need to determine the optimal number of cached items before sending
    // Need to minimize UX impacts
    if (interactionCache.length >= 25) {
        sendData(interactionCache);
        interactionCache = [];
    }
}

export default cacheInteraction;



// Another alternative to this strategy would be using IndexedDB
    // Better for larged amounts of data
    // Async and non-blocking