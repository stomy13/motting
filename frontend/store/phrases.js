import axios from 'axios'

export const state = () => ({
  list: []
})

export const mutations = {
  set(state, phrases) {
    state.list = phrases
  }
}

export const actions = {
  async fetch({ commit }) {
    const res = await axios.get(this.$config.api_base_url + 'phrase')
    const phrases = res.data
    commit('set', phrases)
  },
  async post({ commit }, { text, author }) {
    const params = new URLSearchParams()
    params.append('userid', 'whitebox')
    params.append('text', text)
    params.append('author', author)
    const res = await axios.post(this.$config.api_base_url + 'phrase', params)
    const phrases = res.data
    commit('set', phrases)
  },
  async delete({ commit }, { id }) {
    const params = new URLSearchParams()
    params.append('id', id)
    const res = await axios.delete(this.$config.api_base_url + 'phrase', {
      data: params
    })
    const phrases = res.data
    commit('set', phrases)
  }
}
