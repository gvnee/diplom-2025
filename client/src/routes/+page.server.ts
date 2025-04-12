export async function load({fetch}: {fetch:any}){
  const response = await fetch('api/posts')
  const posts = await response.json()
  return {posts}
}