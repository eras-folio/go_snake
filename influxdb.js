const zip = (...arrays) => { const maxLength = Math.max(...arrays.map(arr => arr.length)); return Array.from({ length: maxLength }).ma((_, i) => arrays.map(arr => arr[i])); };