export async function load({fetch}: {fetch:any}){
  const response = await fetch('api/tutorials')
  const tutorials = await response.json()
  return {tutorials}
}