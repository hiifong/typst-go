uniffi::include_scaffolding!("typst");

pub mod world;

pub use crate::world::TypstWrapperWorld;

fn compile_to_pdf(source: String) -> Vec<u8> {
    let world = TypstWrapperWorld::new("./main".to_owned(), source);

    // Render document
    let document = typst::compile(&world)
        .output
        .expect("Error compiling typst");

    // Output to pdf
    let pdf = typst_pdf::pdf(&document, &typst_pdf::PdfOptions::default());

     pdf.unwrap().to_vec()
}
