/**
 * prettyTime returns a human-friendly time representation from milliseconds.
 * @param {number} ms
 * @returns {(string)[]}
 */
export function prettyTime(ms) {
  const seconds = (ms / 1000).toFixed(1)
  const minutes = (ms / (1000 * 60)).toFixed(1)
  const hours = (ms / (1000 * 60 * 60)).toFixed(1)
  const days = (ms / (1000 * 60 * 60 * 24)).toFixed(1)

  if (ms < 1000) {
    return [ms, '毫秒']
  } else if (seconds < 60) {
    return [seconds, '秒']
  } else if (minutes < 60) {
    return [minutes, '分钟']
  } else if (hours < 24) {
    return [hours, '小时']
  } else {
    return [days, '天']
  }
}
