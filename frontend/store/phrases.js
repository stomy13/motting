import axios from 'axios'

export const state = () => ({
  list: []
})

export const mutations = {
  add(state, phrase) {
    let id = 1
    if (state.list.length > 0) {
      const lastPhrase = state.list[state.list.length - 1]
      id = lastPhrase.ID + 1
    }
    phrase.ID = id
    state.list.push(phrase)
  },
  remove(state, id) {
    const result = state.list.find((phrase) => phrase.ID === id)
    state.list.splice(state.list.indexOf(result), 1)
  },
  set(state, phrases) {
    state.list = phrases
  }
}

export const actions = {
  async fetch({ commit }) {
    const res = await axios.get(this.$config.api_base_url + 'phrase')
    const phrases = res.data
    commit('set', phrases)
  }
}
