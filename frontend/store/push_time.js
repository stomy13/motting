import axios from 'axios'

export const state = () => ({
  push_at: ''
})

export const mutations = {
  modify(state, pt) {
    state.push_at = pt
  },
  reset(state) {
    state.push_at = ''
  }
}

export const actions = {
  async fetch({ commit, state }) {
    const res = await axios.get('http://localhost:3001/api/v1/pushtime')
    const pt = res.data.PushAt
    commit('modify', pt)
  }
}
