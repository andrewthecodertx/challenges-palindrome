# JavaScript

## Solution

```javascript
function isPalindrome(s) {
  const cleaned = s.toLowerCase().replace(/[^a-z0-9]/g, '');
  return cleaned === cleaned.split('').reverse().join('');
}

module.exports = isPalindrome;
```

## Running Tests

```bash
npm install
npm test
```
