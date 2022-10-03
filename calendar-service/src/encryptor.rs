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
        assert_eq!(padded_data.len(), 16, "Padding failed for short string");

        // Test with long string (more than block size)
        let data = "Hello, World! This is a long string.";
        let padded_data = super::pkcs5_padding(data.as_bytes());
        assert_eq!(padded_data.len(), 48, "Padding failed for long string");
    }

    #[test]
    fn test_pkcs5_unpadding() {
        // Test with short string (less than block size)
        let data = "Hello";
        let padded_data: Vec<u8> = super::pkcs5_padding(data.as_bytes());
        let unpadded_data = super::pkcs5_unpadding(&padded_data);
        assert_eq!(unpadded_data.len(), 5, "Unpadding failed for short string");

        // Test with long string (more than block size)
        let data = "Hello, World! This is a long string.";
        let padded_data = super::pkcs5_padding(data.as_bytes());
        let unpadded_data = super::pkcs5_unpadding(&padded_data);
        assert_eq!(unpadded_data.len(), 36, "Unpadding failed for long string");
    }
}
