const get = async <T>(
  url: string,
  options?: RequestInit
): Promise<[T | null, unknown]> => {
  try {
    const response = await fetch(url, {
      headers: { 'Content-Type': 'application/json' },
      ...options
    });
    if (!response.ok)
      throw new Error(`Fetch Error: ${response.status} ${response.statusText}`);
    const json = await response.json();
    return [json as T, null];
  } catch (error) {
    return [null, error];
  }
};

interface Settings {
  header: boolean;
  stringify: boolean;
}

const post = async <T>(
  url: string,
  body: any,
  options?: RequestInit,
  settings: Settings = {
    header: true,
    stringify: true
  }
): Promise<[T | null, unknown]> => {
  try {
    const payload = {
      ...options,
      method: 'POST',
      body: settings.stringify ? JSON.stringify(body) : body
    };
    if (settings.header)
      payload.headers = { 'Content-Type': 'application/json' };
    const response = await fetch(url, payload);
    if (!response.ok)
      throw new Error(`Fetch Error: ${response.status} ${response.statusText}`);
    const json = await response.json();
    return [json as T, null];
  } catch (error) {
    return [null, error];
  }
};

export const fetcher = {
  get,
  post
};
