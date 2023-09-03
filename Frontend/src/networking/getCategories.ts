import { baseUrl } from "./const";

export interface Category {
  Id: number;
  Name: string;
}

export async function getCategories(): Promise<Category[]> {
  const categoriesResponse = await fetch(baseUrl);
  const categories: Promise<Category[]> = await categoriesResponse.json();
  return categories;
}
