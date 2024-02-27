// 1
function gradingStudents(grades) {
  console.log("-- start function --")
  
  // siapkan variable baru untuk menyimpan hasil
  let newGrades = [] 
  for(let i = 0; i < grades.length; i++) {
      // cari kelipatan 5 dari tiap variable input.
      // carbagi angka tersebut dengan 5 dan bulatkan dengan Math.ceil
      // untuk mendapatkan bilangan real lebih besar.
      // lalu kalikan dengan angka 5 untuk mendapatkan kelipatan lima kembali
      let multipleNumber = Math.ceil(grades[i]/5) * 5
      
      // hasil operasi kelipatan 5 kurangi dengan angka actual variable untuk melihat selisih
      let diff = multipleNumber - grades[i]
      console.log(grades[i] + " different " + diff)
      
      // logika grading, jika selisih kurang dari 3 dan angka variable
      // lebih besar dari 38, maka simpan nilai baru
      if (diff < 3 && grades[i] >= 38) {
          newGrades.push(multipleNumber)
      } else {
        if (grades[i] >= 5 ) {
          newGrades.push(grades[i])
        }
      }
  }
  
  console.log("-- finish function --")
  console.log("\n")
  return newGrades
}

let newCollections = new Array(4, 73, 67, 38, 33)
let testGradingStudents = gradingStudents(newCollections)

testGradingStudents.forEach((x, i) => {
  console.log(x)
})


