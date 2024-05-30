import trackClicks from './events/TrackClicks.js'
import trackContentInteraction from './events/TrackContentInteraction.js'
import trackDownloads from './events/TrackDownloads.js'
// import trackErrors from './events/TrackErrors.js'
import trackForms from './events/TrackForms.js'
import trackHovers from './events/TrackHovers.js'
import trackPageResizes from './events/TrackPageResizes.js'
import trackPageTransitions from './events/TrackPageTransitions.js'
import trackScrollDepth from './events/TrackScrollDepth.js'
import trackSession from './events/TrackSession.js'


/**
 * 
 */
function init() {
    trackClicks();
    trackContentInteraction();
    trackDownloads();
    trackForms();
    trackHovers();
    trackPageResizes();
    trackPageTransitions();
    trackScrollDepth();
    trackSession();
}

document.addEventListener('DOMContentLoaded', init);
