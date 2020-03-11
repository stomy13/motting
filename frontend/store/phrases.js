import axios from 'axios'

export const state = () => ({
  list: []
})

export const mutations = {
  set(state, phrases) {
    state.list = phrases
  },
  remove(state, id) {
    const result = state.list.find((phrase) => phrase.ID === id)
    state.list.splice(state.list.indexOf(result), 1)
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
  }
}
