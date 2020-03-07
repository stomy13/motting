export const state = () => ({
  list: [
    {
      text: 'test',
      author: 'test author'
    },
    {
      text: 'test2',
      author: 'test author2'
    },
    {
      text: 'test3',
      author: 'test author3'
    },
    {
      text: 'test4',
      author: 'test author4'
    }
  ]
})

export const mutations = {
  add(state, phrase) {
    state.list.push(phrase)
  },
  remove(state, { phrase }) {
    state.list.splice(state.list.indexOf(phrase), 1)
  }
}
