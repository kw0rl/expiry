export async function checkHealth(): Promise<void> {
  const response = await fetch('/api/health', {
    signal: AbortSignal.timeout(5000),
  })
  if (!response.ok) {
    throw new Error('Unable to reach the API. Check that the Go backend is running.')
  }
  const data: unknown = await response.json()
  if (
    typeof data !== 'object' || data === null ||
    !('status' in data) || data.status !== 'ok' ||
    !('service' in data) || data.service !== 'expiry-api'
  ) {
    throw new Error('The API returned an unexpected response.')
  }
}
