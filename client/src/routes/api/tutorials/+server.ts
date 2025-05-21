import type { Tutorial } from '$lib/types'
import type { Chapter } from '$lib/types'
import {json} from '@sveltejs/kit'

async function getTutorials(){
  const allTutorials = import.meta.glob('/src/tutorials/**/*.md', {eager: true})

  let tutorials: Tutorial[] = []

  for(let path in allTutorials){
    const file = allTutorials[path]
    let slug = path.replace('/src/tutorials/', '').replace('.md', '')
    const tutorialName = slug.split('/')[0]

    if(!tutorials.find(({name}) => name == tutorialName)) {
      tutorials.push({name: tutorialName, chapters: []})
    }

    if(file && typeof file === 'object' && 'metadata' in file){
      const metadata = file.metadata as Omit<Chapter, 'slug'>
      const chapter = {...metadata, slug} satisfies Chapter
      tutorials[tutorials.length - 1].chapters.push(chapter)
    }

  }
  return tutorials
}

export async function GET(){
  const tutorials = await getTutorials()
  return json(tutorials)
}