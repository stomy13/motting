// injectするとvuexでも使える
export default ({ app }, inject) => {
  inject('config', {
    api_base_url: 'http://localhost:3001/api/v1/'
  })
}
