export abstract class TokenService {
  static readonly #key = 'token';
  static readonly #type = 'type';

  static get token() {
    return localStorage.getItem(this.#key);
  }

  static set token(token) {
    console.log('here');
    if (token) localStorage.setItem(this.#key, token);
    else localStorage.removeItem(this.#key);
  }

  static get type() {
    return localStorage.getItem(this.#type);
  }

  static set type(type) {
    if (type) localStorage.setItem(this.#type, type);
    else localStorage.removeItem(this.#type);
  }
}
