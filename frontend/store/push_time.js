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
  async fetch({ commit }) {
    // todo:userid
    const res = await axios.get(this.$config.api_base_url + 'pushtime')
    const pt = res.data.PushAt
    commit('modify', pt)
  },
  async patch({ commit }, pushAt) {
    const params = new URLSearchParams()
    params.append('userid', 'whitebox')
    params.append('pushAt', pushAt)
    const res = await axios.patch(
      this.$config.api_base_url + 'pushtime',
      params
    )
    const pt = res.data.PushAt
    commit('modify', pt)
  }
}
