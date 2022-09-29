export const getSpeciesImagePath = (speciesImageId: string) =>
  `${process.env.NEXT_PUBLIC_API_ENDPOINT}/species-image/${speciesImageId}`
