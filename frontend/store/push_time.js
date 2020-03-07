export const state = () => ({
  push_at: '10:10'
})

export const mutations = {
  modify(state, pt) {
    state.push_at = pt
  },
  reset(state) {
    state.push_at = ''
  }
}
