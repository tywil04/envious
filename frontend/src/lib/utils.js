export function secondsToHHMMSS(totalSeconds, forceHours=false) {
    let hours   = Math.floor(totalSeconds / 3600);
    let minutes = Math.floor((totalSeconds - (hours * 3600)) / 60);
    let seconds = totalSeconds - (hours * 3600) - (minutes * 60);

    if (Number.isNaN(hours) || Number.isNaN(minutes) || Number.isNaN(seconds)) {
        return "00:00"
    }

    let hoursString = hours < 10 ? "0" + hours.toString() : hours.toString()
    let minutesString = minutes < 10 ? "0" + minutes.toString() : minutes.toString()
    let secondsString = seconds < 10 ? "0" + seconds.toString() : seconds.toString()

    let result = minutesString + ":" + secondsString

    if (hoursString !== "00" || forceHours === true) {
        result = hoursString + ":" + result
    }

    return result
}

export function calculateWatchedDuration(watchedSeconds, totalSeconds) {
    watchedSeconds = Math.round(watchedSeconds)
    totalSeconds = Math.round(totalSeconds)

    let totalSecondsHHMMSS = secondsToHHMMSS(totalSeconds)
    let forceHours = totalSecondsHHMMSS.split(":").length === 3 // hours is included
    let watchedSecondsHHMMSS = secondsToHHMMSS(watchedSeconds, forceHours)

    return watchedSecondsHHMMSS + " / " + totalSecondsHHMMSS
}