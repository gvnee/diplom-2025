import {error} from '@sveltejs/kit'
import type { PageLoad } from './$types'

export const load: PageLoad = async({fetch, params}) => {
  try {
    const chapter = await import(/* @vite-ignore */`/src/tutorials/${params.tutorial}/${params.chapterSlug}.md`)
    // await fetch('http://localhost:4000/test', {
    //   method: 'POST',
    //   body: JSON.stringify(chapter.metadata)
    // }).then((response) => {return response.text()})
    return {
      content: chapter.default,
      meta: chapter.metadata
    }
  } catch(e){
    error(404, `Could not find ${params.tutorial}/${params.chapterSlug}`)
  }
}