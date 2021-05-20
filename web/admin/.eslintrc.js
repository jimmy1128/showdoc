module.exports = {
  root: true,
  env: {
    node: true
  },
  extends: [
    'plugin:vue/essential',
    '@vue/standard'
  ],
  parserOptions: {
    parser: 'babel-eslint'
  },
  rules: {
    'no-console': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
    'no-debugger': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
    'vue/require-prop-type-constructor': 0,
    'no-callback-literal': 0,
    'one-var': [2, { "initialized": "never" }],
    'semi-spacing': 0,
    'no-void': 0,
    'no-sequences': 2,
    camelcase: ['warn', { properties: 'never' }]
  }
}
