export async function GetVerison() {
    const response = await fetch("http://localhost:8080/version") 

    if (!response.ok) {
        const error = await response.json().catch(() => null)
        throw new Error(error?.error ?? "Failed to get version.")
    }

    return response.json() as Promise<{version: string}>
}