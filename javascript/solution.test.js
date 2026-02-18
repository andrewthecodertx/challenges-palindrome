const isPalindrome = require('./solution');

describe('Palindrome', () => {
  test('simple palindromes', () => {
    expect(isPalindrome('racecar')).toBe(true);
    expect(isPalindrome('hello')).toBe(false);
  });
  test('with spaces and punctuation', () => {
    expect(isPalindrome('A man, a plan, a canal: Panama')).toBe(true);
  });
  test('empty string', () => {
    expect(isPalindrome('')).toBe(true);
  });
});
