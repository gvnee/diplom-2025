import {error} from '@sveltejs/kit'

export async function load({params}: {params:any}){
  try {
    const post = await import(`../../posts/${params.slug}.md`)

    return {
      content: post.default,
      meta: post.meta
    }
  } catch(e){
    error(404, `Could not find ${params.slug}`)
  }
}