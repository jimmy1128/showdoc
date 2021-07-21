module.exports = {
  root: true,
  env: {
    browser: true
  },
  extends: ["standard", "plugin:vue/recommended"],
  parser: 'vue-eslint-parser',
  parserOptions: {
    sourceType: 'module'
  },
  plugins: [
    'html'
  ],
  rules: {
    // 'no-debugger': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
    'vue/require-prop-type-constructor': 0,
    'no-debugger': process.env.NODE_ENV === 'production' ? 'error' : 'off',
    'no-callback-literal': 0,
    'one-var': [2, { "initialized": "never" }],
    'semi-spacing': 0,
    'no-void': 0,
    'no-sequences': 2,
    'camelcase': 'off',
    camelcase: ['warn', { properties: 'never' }]
  }
}
