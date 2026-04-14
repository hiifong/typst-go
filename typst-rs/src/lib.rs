uniffi::include_scaffolding!("typst");

pub mod world;

use typst::{layout::Abs, visualize::Color};

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

fn compile_to_svg(source: String) -> Vec<u8> {
    let world = TypstWrapperWorld::new("./main".to_owned(), source);
    let document = typst::compile(&world).output.expect("Error compliling typst");

    let merged = typst_svg::svg_merged(&document, Abs::pt(14.0));
    
    merged.as_bytes().to_vec()
}

fn compile_to_png(source: String) -> Vec<u8>{
    let world = TypstWrapperWorld::new("./main".to_owned(), source);
    let document = typst::compile(&world).output.expect("Error compliling typst");

    let merged = typst_render::render_merged(&document, 2.0, Abs::pt(14.0), Some(Color::WHITE));
    
    merged.encode_png().unwrap().to_vec()
}