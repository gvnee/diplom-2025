import {error} from '@sveltejs/kit'

export async function load({params}: {params:any}){
  try {
    const chapter = await import(`/src/tutorials/${params.tutorial}/${params.chapterSlug}.md`)

    return {
      content: chapter.default,
      meta: chapter.metadata
    }
  } catch(e){
    error(404, `Could not find ${params.tutorial}/${params.slug}`)
  }
}