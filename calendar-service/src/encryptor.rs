use ::generic_array::GenericArray;
use aes_gcm::{
    aead::{AeadInPlace, KeyInit},
    Aes256Gcm, Nonce,
};

const BLOCK_SIZE: usize = 16;

/// It takes a string, encrypts it, and returns the encrypted data as a vector of bytes
///
/// Arguments:
///
/// * `data`: The data to be encrypted.
/// * `key`: AES-256 key encryption key.
/// * `iv`: Initialization vector. This is a random value that is used to ensure that the same plaintext
/// will not produce the same ciphertext.
///
/// Returns:
///
/// A vector of bytes
pub fn encrypt(
    data: String,
    key: &[u8; 32],
    iv: &[u8; 12],
) -> Result<Vec<u8>, Box<dyn std::error::Error>> {
    let cipher = Aes256Gcm::new(GenericArray::from_slice(key));
    let nonce = Nonce::from_slice(iv); // 96-bits; unique per message

    let mut buffer = pkcs5_padding(data.as_bytes());

    // Encrypt `buffer` in-place, replacing the plaintext contents with ciphertext
    cipher
        .encrypt_in_place(nonce, b"", &mut buffer)
        .expect("encryption failure!");

    // `buffer` now contains the message ciphertext
    assert_ne!(&buffer, data.as_bytes());

    Ok(buffer)
}

/// It takes a byte array, decrypts it using the key and iv, and returns a string
///
/// Arguments:
///
/// * `data`: The data to be decrypted.
/// * `key`: AES-256 key encryption key.
/// * `iv`: Initialization vector. This is a random value that is used to ensure that the same plaintext
/// will not produce the same ciphertext.
///
/// Returns:
///
/// A String
pub fn decrypt(data: &[u8], key: &[u8; 32], iv: &[u8; 12]) -> String {
    // TODO: Return Result like in encrypt()
    let cipher = Aes256Gcm::new(GenericArray::from_slice(key));
    let nonce = Nonce::from_slice(iv);

    let mut buffer = data.to_vec();
    cipher
        .decrypt_in_place(nonce, b"", &mut buffer)
        .expect("decryption failure!");

    let buffer = pkcs5_unpadding(&buffer);
    std::str::from_utf8(&buffer)
        .expect("u8 to String transfromation failed")
        .to_string()
}

/// Add padding bytes for the message to make it divisible by the block size
/// before encryption.
///
/// Arguments:
///
/// * `src`: The source data to be padded.
///
/// Returns:
///
/// A vector of bytes.
fn pkcs5_padding(src: &[u8]) -> Vec<u8> {
    let padding = BLOCK_SIZE - src.len() % BLOCK_SIZE;
    let padtext = vec![padding as u8; padding];
    let mut result = src.to_vec();
    result.append(&mut padtext.to_vec());
    return result;
}

/// Remove padding bytes from the end of the slice (used after decryption)
///
/// Arguments:
///
/// * `src`: Decrypted source data.
///
/// Returns:
///
/// A vector of bytes.
fn pkcs5_unpadding(src: &[u8]) -> Vec<u8> {
    let length = src.len();
    let unpadding = src[length - 1] as usize;
    if length < unpadding {
        panic!(
            "Invalid padding. length: {}, unpadding: {}",
            length, unpadding
        );
    }
    (&src[..(length - unpadding)]).to_vec()
}

#[cfg(test)]
mod tests {
    #[test]
    fn test_pkcs5_padding() {
        // Test with short string (less than block size)
        let data = "Hello";
        let padded_data = super::pkcs5_padding(data.as_bytes());
        let expected: Vec<u8> = vec![
            72, 101, 108, 108, 111, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
        ];
        assert_eq!(padded_data, expected, "Padding failed for short string");

        // Test with long string (more than block size)
        let data = "Hello, Big World!";
        let padded_data = super::pkcs5_padding(data.as_bytes());
        let expected: Vec<u8> = vec![
            72, 101, 108, 108, 111, 44, 32, 66, 105, 103, 32, 87, 111, 114, 108, 100, 33, 15, 15,
            15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
        ];
        assert_eq!(padded_data, expected, "Padding failed for long string");
    }

    #[test]
    fn test_pkcs5_unpadding() {
        // Test with short string (less than block size)
        let padded_data: Vec<u8> = vec![
            72, 101, 108, 108, 111, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
        ];
        let unpadded_data = super::pkcs5_unpadding(&padded_data);
        let expected: Vec<u8> = vec![72, 101, 108, 108, 111];
        assert_eq!(unpadded_data, expected, "Unpadding failed for short string");

        // Test with long string (more than block size)
        let padded_data: Vec<u8> = vec![
            72, 101, 108, 108, 111, 44, 32, 66, 105, 103, 32, 87, 111, 114, 108, 100, 33, 15, 15,
            15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
        ];
        let unpadded_data = super::pkcs5_unpadding(&padded_data);
        let expected: Vec<u8> = vec![
            72, 101, 108, 108, 111, 44, 32, 66, 105, 103, 32, 87, 111, 114, 108, 100, 33,
        ];
        assert_eq!(unpadded_data, expected, "Unpadding failed for long string");
    }

    const KEY: &[u8; 32] = b"01234567890123456789012345678901";
    const IV: &[u8; 12] = b"012345678901";

    #[test]
    fn test_encrypt() {
        let data = "Hello";
        let encrypted_data = super::encrypt(data.to_string(), KEY, IV).unwrap();
        let expected_data = [
            56, 234, 164, 6, 206, 254, 84, 123, 168, 97, 157, 67, 127, 222, 121, 111, 3, 34, 21,
            236, 175, 184, 216, 246, 118, 243, 15, 247, 102, 209, 63, 138,
        ]
        .to_vec();
        assert_eq!(encrypted_data, expected_data, "Encryption failed");
    }

    #[test]
    fn test_decrypt() {
        let encrypted_data = [
            56, 234, 164, 6, 206, 254, 84, 123, 168, 97, 157, 67, 127, 222, 121, 111, 3, 34, 21,
            236, 175, 184, 216, 246, 118, 243, 15, 247, 102, 209, 63, 138,
        ];
        let decrypted_data = super::decrypt(&encrypted_data, KEY, IV);
        let expected_data = "Hello".to_string();
        assert_eq!(decrypted_data, expected_data, "Decryption failed");
    }
}
