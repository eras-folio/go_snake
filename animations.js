const findLongestWord = (str) => { const words = str.split(’ ‘); return words.reduce((longest, current) => { return current.length > longest.length ? current : longest; }, ‘’); }; console.log(findLongestWord(‘The quick brown fox jumped over the lazy dog’));