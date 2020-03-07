export const state = () => ({
  list: [
    {
      id: 1,
      text: 'test',
      author: 'test author'
    },
    {
      id: 2,
      text: 'test2',
      author: 'test author2'
    },
    {
      id: 3,
      text: 'test3',
      author: 'test author3'
    },
    {
      id: 4,
      text: 'test4',
      author: 'test author4'
    },
    {
      id: 5,
      text: 'test2',
      author: 'test author2'
    },
    {
      id: 6,
      text: 'test3',
      author: 'test author3'
    },
    {
      id: 7,
      text: 'test4',
      author: 'test author4'
    },
    {
      id: 8,
      text: 'test2',
      author: 'test author2'
    },
    {
      id: 9,
      text: 'test3',
      author: 'test author3'
    },
    {
      id: 10,
      text: 'test4',
      author: 'test author4'
    },
    {
      id: 11,
      text: 'test2',
      author: 'test author2'
    },
    {
      id: 12,
      text: 'test3',
      author: 'test author3'
    },
    {
      id: 13,
      text: 'test4',
      author: 'test author4'
    }
  ]
})

export const mutations = {
  add(state, phrase) {
    let id = 1
    if (state.list.length > 0) {
      const lastPhrase = state.list[state.list.length - 1]
      id = lastPhrase.id + 1
    }
    phrase.id = id
    state.list.push(phrase)
  },
  remove(state, id) {
    const result = state.list.find((phrase) => phrase.id === id)
    state.list.splice(state.list.indexOf(result), 1)
  }
}
