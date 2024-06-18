/**
 * Remake of the throttle function from loadash import
 * 
 * @param {function} func 
 * @param {number} wait 
 * @returns 
 */

export function throttle(func, wait = 0) {
    let waiting = false;

    return (...args) => {
      if (!waiting) {
        func(...args);
        waiting = true;
        
        setTimeout(function () {
          waiting = false;
        }, wait);
      }
    };
  }